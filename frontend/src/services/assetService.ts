import { axiosInstance, BASE_URL } from "@/instances/axiosIntance";
import { Asset } from "@/types/asset";
import { CreditStatus } from "@/types/creditStatus";

const assetStatusUrl = BASE_URL + "assets";

export const fetchAssets = async (): Promise<Asset[]> => {
    const response = await axiosInstance.get<Asset[]>(assetStatusUrl);
    return response.data;
};