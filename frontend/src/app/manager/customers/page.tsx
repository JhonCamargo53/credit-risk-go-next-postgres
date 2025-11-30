'use client';

import CustomerManager from "@/components/modules/customer/CustomerManager";
import { CreditRequestProvider } from "@/context/CreditRequestContext";
import { CustomerAssetProvider } from "@/context/CustomerAssetContext";
import { CustomerProvider } from "@/context/CustomerContext";
import ProtectedRoute from "@/hoc/ProtectedRoute";

const CustomerPage = () => {
     return (
          <CustomerProvider >
               <CreditRequestProvider>
                    <CustomerAssetProvider>
                         <CustomerManager />
                    </CustomerAssetProvider>
               </CreditRequestProvider>
          </CustomerProvider>
     )
}

export default ProtectedRoute(CustomerPage)

