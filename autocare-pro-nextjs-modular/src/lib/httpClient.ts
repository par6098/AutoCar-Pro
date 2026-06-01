import { API_BASE_URL } from '@/config/api';

function getAuthHeaders(): HeadersInit {
  const headers: HeadersInit = {
    'Content-Type': 'application/json',
  };

  if (typeof window !== 'undefined') {
    const access_token = localStorage.getItem('access_token');

    if (access_token) {
      headers.Authorization = `Bearer ${access_token}`;
    }
  }

  return headers;
}

export async function httpGet<T>(
  path: string
): Promise<T> {
  const response = await fetch(
    `${API_BASE_URL}${path}`,
    {
      method: 'GET',
      headers: getAuthHeaders(),
      cache: 'no-store',
    }
  );

  if (!response.ok) {
    throw new Error(
      `GET ${path} failed: ${response.status}`
    );
  }

  return response.json() as Promise<T>;
}

export async function httpPost<
  TRequest,
  TResponse
>(
  path: string,
  body: TRequest
): Promise<TResponse> {
  const response = await fetch(
    `${API_BASE_URL}${path}`,
    {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify(body),
    }
  );

  if (!response.ok) {
    throw new Error(
      `POST ${path} failed: ${response.status}`
    );
  }

  return response.json() as Promise<TResponse>;
}

export async function httpPatch<
  TRequest,
  TResponse
>(
  path: string,
  body: TRequest
): Promise<TResponse> {
  const response = await fetch(
    `${API_BASE_URL}${path}`,
    {
      method: 'PATCH',
      headers: getAuthHeaders(),
      body: JSON.stringify(body),
    }
  );

  if (!response.ok) {
    throw new Error(
      `PATCH ${path} failed: ${response.status}`
    );
  }

  return response.json() as Promise<TResponse>;
}