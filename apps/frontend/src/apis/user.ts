import { Key } from '@repo/models';
import apiClient from './apiClient';

export const fetchKeysData = async () => {
  const response = await apiClient.get('api/keys');
  return response;
};

export const updateKeyData = async (payload: Key) => {
  const response = await apiClient.put(`api/keys/${payload.id}`, {
    ...payload
  });
  return response;
};

export const fetchKeyData = async (id: string) => {
  const response = await apiClient.get(`api/keys/${id}`);
  return response;
};
