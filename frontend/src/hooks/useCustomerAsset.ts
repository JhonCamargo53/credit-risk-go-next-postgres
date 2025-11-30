import { CustomerAssetContext, CustomerAssetContextType } from "@/context/CustomerAssetContext";
import { useContext } from "react";

export const useCustomerAsset = (): CustomerAssetContextType => {
    const context = useContext(CustomerAssetContext);
    if (!context) {
        throw new Error('useCustomerAsset debe usarse dentro de un CustomerAssetProvider');
    }
    return context;
};
