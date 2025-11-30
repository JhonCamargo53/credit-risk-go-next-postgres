'use client';

import LoadingPage from "@/components/common/loading/LoadingPage";
import { useAuth } from "@/hooks/useAuth";


export default function LayoutWithAuth({ children }: { children: React.ReactNode }) {
    
    const { loading } = useAuth();

    if (loading) {
        return <LoadingPage loadingText="Cargando sistema"/>
    }

    return <>{children}</>;
}
