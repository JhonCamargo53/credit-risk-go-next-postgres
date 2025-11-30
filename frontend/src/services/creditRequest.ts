import { axiosInstance, BASE_URL } from "@/instances/axiosIntance";

const creditRequestUrl = BASE_URL + "credit-requests";

export const fetchCreditRequests = async (): Promise<CreditRequest[]> => {
    const response = await axiosInstance.get<CreditRequest[]>(creditRequestUrl);
    return response.data;
};

export const fetchCreditRequestById = async (id: number): Promise<CreditRequest> => {
    const response = await axiosInstance.get<CreditRequest>(`${creditRequestUrl}/${id}`);
    return response.data;
};

export const createCreditRequest = async (data: CreditRequestForm): Promise<CreditRequest> => {
    const response = await axiosInstance.post<CreditRequest>(creditRequestUrl, data);
    return response.data;
};

export const updateCreditRequest = async (id: number, data: Partial<CreditRequestForm>): Promise<CreditRequest> => {
    const response = await axiosInstance.put<CreditRequest>(`${creditRequestUrl}/${id}`, data);
    return response.data;
};

export const deleteCreditRequest = async (id: number): Promise<{ message: string }> => {
    const response = await axiosInstance.delete<{ message: string }>(`${creditRequestUrl}/${id}`);
    return response.data;
};

export const fetchCreditRequestsByCustomerId = async (customerId: number): Promise<CreditRequest[]> => {
    const response = await axiosInstance.get<CreditRequest[]>(`${creditRequestUrl}`, { params: { customerId } });
    return response.data;
};
