/**
 * API 密钥管理
 */

import type { ConfigApiKeyEntry } from '@/types';
import { apiClient } from './client';

function normalizeApiKeyEntry(item: unknown): ConfigApiKeyEntry | null {
  if (typeof item === 'string') {
    const key = item.trim();
    return key ? { key } : null;
  }
  if (item === null || typeof item !== 'object' || Array.isArray(item)) return null;
  const record = item as Record<string, unknown>;
  const rawKey = record.key ?? record['api-key'] ?? record.apiKey ?? record.Key;
  const key = String(rawKey ?? '').trim();
  if (!key) return null;
  const remarkRaw = record.remark ?? record.owner;
  const remark = String(remarkRaw ?? '').trim();
  return remark ? { key, remark } : { key };
}

export const apiKeysApi = {
  async list(): Promise<ConfigApiKeyEntry[]> {
    const data = await apiClient.get<Record<string, unknown>>('/api-keys');
    const keys = data['api-keys'] ?? data.apiKeys;
    return Array.isArray(keys)
      ? keys
          .map((item) => normalizeApiKeyEntry(item))
          .filter((item): item is ConfigApiKeyEntry => Boolean(item))
      : [];
  },

  replace: (keys: ConfigApiKeyEntry[]) => apiClient.put('/api-keys', keys),

  update: (index: number, value: ConfigApiKeyEntry) => apiClient.patch('/api-keys', { index, value }),

  delete: (index: number) => apiClient.delete(`/api-keys?index=${index}`)
};
