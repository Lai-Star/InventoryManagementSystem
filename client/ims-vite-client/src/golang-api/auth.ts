import axios, { AxiosResponse } from 'axios';

type LoginResponse = {
  Success: string;
  Status: number;
};

type LoginPayload = {
  username: string;
  password: string;
};

interface SignUpPayload {
  username: string;
  password: string;
  email: string;
}

interface SignUpResponse {
  Success: string;
  Status: number;
}

export const loginRoute = async (
  username: string,
  password: string
): Promise<AxiosResponse<LoginResponse>> => {
  const response = await axios.post<LoginPayload, AxiosResponse<LoginResponse>>(
    'http://localhost:8080/login',
    {
      username,
      password,
    }
  );
  return response;
};

export const signUpRoute = async (
  username: string,
  password: string,
  email: string
): Promise<AxiosResponse<SignUpResponse>> => {
  const response = await axios.post<
    SignUpPayload,
    AxiosResponse<SignUpResponse>
  >('http://localhost:8080/signup', {
    username,
    password,
    email,
  });

  return response;
};