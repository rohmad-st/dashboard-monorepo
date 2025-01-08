import ApiClient from './apiClient';

const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8000';

const apiClient = new ApiClient(API_BASE_URL);

export default apiClient;
