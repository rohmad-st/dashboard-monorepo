import { fetchKeyData, fetchKeysData, updateKeyData } from '@/apis/user';
import { createAsyncThunk } from '@reduxjs/toolkit';

// get list of users
export const getUsersData = createAsyncThunk(
  'user/getUsersData',
  fetchKeysData
);
// update user
export const putUserData = createAsyncThunk(
  'user/updateKeyData',
  updateKeyData
);

// get single user by id
export const getUserData = createAsyncThunk('user/getUserData', fetchKeyData);
