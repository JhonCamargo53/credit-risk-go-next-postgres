import { CustomerContext, CustomerContextType } from "@/context/CustomerContext";
import { useContext } from "react";

export const useCustomer = (): CustomerContextType => {
    const context = useContext(CustomerContext);
    if (!context) {
        throw new Error('useCustomer debe usarse dentro de un CustomerProvider');
    }
    return context;
};
