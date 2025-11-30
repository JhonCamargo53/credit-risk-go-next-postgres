import { axiosInstance, BASE_URL } from "@/instances/axiosIntance";
import { CustomerAsset, CustomerAssetForm } from "@/types/customerAsset";

const managementUrl = BASE_URL + "customer-assets";

export const fetchCustomerAssets = async (): Promise<CustomerAsset[]> => {
    const response = await axiosInstance.get<CustomerAsset[]>(managementUrl);
    return response.data;
};

export const fetchCustomerAssetById = async (id: number): Promise<CustomerAsset> => {
    const response = await axiosInstance.get<CustomerAsset>(`${managementUrl}/${id}`);
    return response.data;
};

export const fetchCustomerAssetsByCreditRequest = async (id: number): Promise<CustomerAsset[]> => {
    const response = await axiosInstance.get<CustomerAsset[]>(`${managementUrl}`, { params: { creditRequestId: id } });
    return response.data;
};

export const createCustomerAsset = async (data: CustomerAssetForm): Promise<CustomerAsset> => {
    const response = await axiosInstance.post<CustomerAsset>(managementUrl, data);
    return response.data;
};

export const updateCustomerAsset = async (
    id: number,
    data: Partial<Omit<CustomerAsset, "ID" | "status">>
): Promise<CustomerAsset> => {
    const response = await axiosInstance.put<CustomerAsset>(`${managementUrl}/${id}`, data);
    return response.data;
};

export const deleteCustomerAsset = async (id: number): Promise<{ message: string }> => {
    const response = await axiosInstance.delete<{ message: string }>(`${managementUrl}/${id}`);
    return response.data;
};
