import axios, { AxiosResponse } from 'axios';

type LoginData = {
  Success: string;
  Status: number;
};

type LoginPayload = {
  username: string;
  password: string;
};

const loginRoute = async (
  username: string,
  password: string
): Promise<AxiosResponse<LoginData>> => {
  const response = await axios.post<LoginPayload, AxiosResponse<LoginData>>(
    'http://localhost:8080/login',
    {
      username,
      password,
    }
  );
  return response;
};

export default loginRoute;
