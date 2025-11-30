interface CreditRequest {
    ID: number
    amount: number;
    termMonths: number;
    productType: string
    creditStatusId: number
    riskScore: number
    riskCategory: string
    riskExplanation: string
    customerId: number;
    UpdatedAt: string;
    CreatedAt: string;
}

interface CreditRequestForm {
    amount: number;
    termMonths: number;
    productType: string;
    creditStatusId: number;
    customerId:number;
}