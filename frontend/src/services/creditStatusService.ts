import { axiosInstance, BASE_URL } from "@/instances/axiosIntance";
import { CreditStatus } from "@/types/creditStatus";

const creditStatusUrl = BASE_URL + "credit-statuses";

export const fetchCreditStatuses = async (): Promise<CreditStatus[]> => {
    const response = await axiosInstance.get<CreditStatus[]>(creditStatusUrl);
    return response.data;
};