'use client';
import React from "react";
import { useAuth } from "@/hooks/useAuth";
import { AvailableRoles, roleDictionary } from "@/config/role.config";
import LoadingPage from "@/components/common/loading/LoadingPage";
import PageNotFound from "@/components/common/utils/PageNotFound";
import { usePathname } from "next/navigation";
import { getAllowedRolesByPath } from "@/utils/routeUtils";

export default function ProtectedRoute<P extends object>(
    Component: React.ComponentType<P>,
     allowedRoles?: AvailableRoles[]
) {
    const ProtectedWrapper = (props: P) => {
       
        const { user, loading } = useAuth();

        const pathname = usePathname().replace('manager/','');
        if (loading) {
            return <LoadingPage loadingText="Validando sesión..." />;
        }

        if (!user) {
            return <PageNotFound />;
        }

        const finalAllowedRoles: AvailableRoles[] = [...getAllowedRolesByPath(pathname)];
        const canAccess = finalAllowedRoles.some(
            (allowedRoleItem) => roleDictionary[allowedRoleItem] === user.roleId
        );

        return canAccess ? <Component {...props} /> : <PageNotFound />;
    };

    ProtectedWrapper.displayName = `ProtectedRoute(${Component.displayName || Component.name || 'Component'})`;

    return ProtectedWrapper;
}
