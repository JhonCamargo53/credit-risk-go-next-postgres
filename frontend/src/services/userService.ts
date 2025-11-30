import { axiosInstance, BASE_URL } from "@/instances/axiosIntance";
import { User, UserForm } from "@/types/user";

const managementUrl = BASE_URL + "users";

export const fetchUsers = async (): Promise<User[]> => {
    const response = await axiosInstance.get<User[]>(managementUrl);
    return response.data;
};

export const fetchUserById = async (id: number): Promise<User> => {
    const response = await axiosInstance.get<User>(`${managementUrl}/${id}`);
    return response.data;
};

export const createUser = async (data: UserForm): Promise<User> => {
    const response = await axiosInstance.post<User>(managementUrl, data);
    return response.data;
};

export const updateUser = async (id: number, data: Partial<Omit<User, "ID"| "status">>): Promise<User> => {
    const response = await axiosInstance.put<User>(`${managementUrl}/${id}`, data);
    return response.data;
};

export const deleteUser = async (id: number): Promise<{ message: string }> => {
    const response = await axiosInstance.delete<{ message: string }>(`${managementUrl}/${id}`);
    return response.data;
};
