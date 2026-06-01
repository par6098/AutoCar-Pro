import { USE_MOCKS } from '@/config/api';
import { httpGet } from '@/lib/httpClient';
import { MessageThread } from '@/types/domain';
import { messageThreads } from './mockData';
export const messageService = { listThreads: () => USE_MOCKS ? Promise.resolve(messageThreads) : httpGet<MessageThread[]>('/messages/threads') };
