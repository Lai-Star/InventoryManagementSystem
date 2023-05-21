import axios, { AxiosResponse } from 'axios';

export const getAccountsRoute = async () => {
  const response = await axios.get('http://localhost:8080/admin/users', {
    withCredentials: true, // Include credentials (cookies) in the request
  });

  return response;
};
