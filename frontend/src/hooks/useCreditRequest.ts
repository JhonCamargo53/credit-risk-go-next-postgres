import { CreditRequestContext, CreditRequestContextType } from "@/context/CreditRequestContext";
import { useContext } from "react";

export const useCreditRequest = (): CreditRequestContextType => {
    const context = useContext(CreditRequestContext);
    if (!context) {
        throw new Error('useCreditRequest debe usarse dentro de un CreditRequestProvider');
    }
    return context;
};
