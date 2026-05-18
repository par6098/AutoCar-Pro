package internal

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	db  *pgxpool.Pool
	cfg Config
}

func NewHandler(db *pgxpool.Pool, cfg Config) *Handler {
	return &Handler{db: db, cfg: cfg}
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var req RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if req.Role == "" {
		req.Role = "customer"
	}

	if req.Role != "admin" && req.Role != "employee" && req.Role != "customer" {
		return c.Status(400).JSON(fiber.Map{"error": "invalid role"})
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "password hashing failed"})
	}

	userID := uuid.New().String()

	_, err = h.db.Exec(context.Background(),
		`INSERT INTO users (id, email, password_hash, role, status, created_at)
		 VALUES ($1, $2, $3, $4, $5, NOW())`,
		userID,
		req.Email,
		string(passwordHash),
		req.Role,
		"active",
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "user creation failed"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "user registered successfully",
		"user_id": userID,
	})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	var user User

	err := h.db.QueryRow(context.Background(),
		`SELECT id, email, password_hash, role, status, created_at
		 FROM users
		 WHERE email = $1`,
		req.Email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.Status,
		&user.CreatedAt,
	)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	if user.Status != "active" {
		return c.Status(403).JSON(fiber.Map{"error": "user inactive"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	accessToken, err := GenerateAccessToken(user, h.cfg)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "access token generation failed"})
	}

	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "refresh token generation failed"})
	}

	refreshHash := hashToken(refreshToken)
	expiresAt := time.Now().Add(time.Duration(h.cfg.RefreshTokenDays) * 24 * time.Hour)

	_, err = h.db.Exec(context.Background(),
		`INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at, revoked, created_at)
		 VALUES ($1, $2, $3, $4, false, NOW())`,
		uuid.New().String(),
		user.ID,
		refreshHash,
		expiresAt,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "refresh token save failed"})
	}

	return c.JSON(TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	})
}

func (h *Handler) RefreshToken(c *fiber.Ctx) error {
	var req RefreshRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	refreshHash := hashToken(req.RefreshToken)

	var user User
	var tokenID string

	err := h.db.QueryRow(context.Background(),
		`SELECT rt.id, u.id, u.email, u.password_hash, u.role, u.status, u.created_at
		 FROM refresh_tokens rt
		 JOIN users u ON u.id = rt.user_id
		 WHERE rt.token_hash = $1
		   AND rt.revoked = false
		   AND rt.expires_at > NOW()`,
		refreshHash,
	).Scan(
		&tokenID,
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.Status,
		&user.CreatedAt,
	)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid refresh token"})
	}

	_, _ = h.db.Exec(context.Background(),
		`UPDATE refresh_tokens SET revoked = true WHERE id = $1`,
		tokenID,
	)

	newAccessToken, err := GenerateAccessToken(user, h.cfg)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "access token generation failed"})
	}

	newRefreshToken, err := GenerateRefreshToken()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "refresh token generation failed"})
	}

	newRefreshHash := hashToken(newRefreshToken)
	expiresAt := time.Now().Add(time.Duration(h.cfg.RefreshTokenDays) * 24 * time.Hour)

	_, err = h.db.Exec(context.Background(),
		`INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at, revoked, created_at)
		 VALUES ($1, $2, $3, $4, false, NOW())`,
		uuid.New().String(),
		user.ID,
		newRefreshHash,
		expiresAt,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "refresh token save failed"})
	}

	return c.JSON(TokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		TokenType:    "Bearer",
	})
}

func (h *Handler) Me(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"user_id": c.Locals("user_id"),
		"email":   c.Locals("email"),
		"role":    c.Locals("role"),
	})
}

func (h *Handler) AdminDashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "admin dashboard"})
}

func (h *Handler) EmployeeDashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "employee dashboard"})
}

func (h *Handler) CustomerDashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "customer dashboard"})
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
