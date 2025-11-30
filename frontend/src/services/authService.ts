import { axiosInstance, BASE_URL } from "@/instances/axiosIntance";
import { LoginForm } from "@/types/auth";

const managementUrl = BASE_URL

export const signIn = async (formData: LoginForm) => {
    const response =  await axiosInstance.post(managementUrl + 'login', { ...formData })
    return response.data
}
