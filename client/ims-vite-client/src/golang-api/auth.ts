import axios, { AxiosResponse } from 'axios';

const loginRoute = async (
  username: string,
  password: string
): Promise<AxiosResponse> => {
  return await axios.post('http://localhost:8080/login', {
    username,
    password,
  });
};

export default loginRoute;
