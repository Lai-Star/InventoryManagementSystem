import axios, { AxiosResponse } from 'axios';

type GetAccountsData = {
  UserId: number;
  Username: string;
  Email: string;
  IsActive: number;
  OrganisationName: string;
};

interface GetAccountsResponse {
    Result: GetAccountsData[];
    Status: number;
    Success: string;
}

export const getAccountsRoute = async (): Promise<
  AxiosResponse<GetAccountsResponse>
> => {
  const response = await axios.get<GetAccountsResponse>(
    'http://localhost:8080/admin/users',
    {
      withCredentials: true, // Include credentials (cookies) in the request
    }
  );

  return response;
};
