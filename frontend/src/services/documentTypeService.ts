import { axiosInstance, BASE_URL } from "@/instances/axiosIntance";
import { Customer, CustomerForm } from "@/types/customer";
import { DocumentType } from "@/types/documentType";

const managementUrl = BASE_URL + "document-types";

export const fetchDocumentTypes = async (): Promise<DocumentType[]> => {
    const response = await axiosInstance.get<DocumentType[]>(managementUrl);
    return response.data;
};