export interface Customer {
    ID: number;
    name: string;
    email: string;
    phoneNumber: string;
    documentNumber: string;
    documentTypeId: number;
    monthlyIncome: number;
    createdById: number;
    status: boolean;
    UpdatedAt: string;
    CreatedAt: string;
}

export interface CustomerForm {
    name: string;
    email: string;
    phoneNumber: string;
    documentNumber: string;
    documentTypeId: number;
    monthlyIncome: number;
    status: boolean;
}