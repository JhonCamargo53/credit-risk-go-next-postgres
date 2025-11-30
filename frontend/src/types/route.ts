import { AvailableRoles } from "@/config/role.config"
import { JSX } from "react";

export interface IRoute {
    label: string;
    path: string;
    allowedRoles: AvailableRoles[];
    icon: JSX.Element;
    renderNavigator: boolean;
    subpaths?: IRoute[];
}