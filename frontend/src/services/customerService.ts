import { axiosInstance, BASE_URL } from "@/instances/axiosIntance";
import { Customer, CustomerForm } from "@/types/customer";

const managementUrl = BASE_URL + "customers";

export const fetchCustomers = async (): Promise<Customer[]> => {
    const response = await axiosInstance.get<Customer[]>(managementUrl);
    return response.data;
};

export const fetchCustomerById = async (id: number): Promise<Customer> => {
    const response = await axiosInstance.get<Customer>(`${managementUrl}/${id}`);
    return response.data;
};

export const createCustomer = async (
    data: CustomerForm
): Promise<Customer> => {
    const response = await axiosInstance.post<Customer>(managementUrl, data);
    return response.data;
};

export const updateCustomer = async (
    id: number,
    data: Partial<Omit<Customer, "ID" | "status">>
): Promise<Customer> => {
    const response = await axiosInstance.put<Customer>(`${managementUrl}/${id}`, data);
    return response.data;
};

export const deleteCustomer = async (id: number): Promise<{ message: string }> => {
    const response = await axiosInstance.delete<{ message: string }>(`${managementUrl}/${id}`);
    return response.data;
};
