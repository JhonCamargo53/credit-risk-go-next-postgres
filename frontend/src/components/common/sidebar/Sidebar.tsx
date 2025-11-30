'use client';
import { useState } from "react";
import Link from "next/link";
import { routeAccessConfig } from "@/config/route.access.config";
import { IoMdArrowDropup } from "react-icons/io";
import React from 'react';
import { useAuth } from "@/hooks/useAuth";
import { roleDictionary } from "@/config/role.config";
import { usePathname } from "next/navigation";
import { IRoute } from "@/types/route";

interface SidebarProps {
    open: boolean;
}

const Sidebar: React.FC<SidebarProps> = ({ open }) => {
    return (
        <div
            className="bg-neutral-dark overflow-hidden transition-transform duration-300 ease-in-out left-0 h-full fixed mt-0 px-3 z-50"
            style={{
                transform: open ? 'translateX(0)' : 'translateX(-100%)',
                width: '16rem',
            }}
        >
            {routeAccessConfig.map((route, index) => (
                <SidebarItem key={index} route={route} />
            ))}
        </div>
    );
};

export default Sidebar;

function SidebarItem({ route }: { route: IRoute }) {
    const { user } = useAuth();
    const [open, setOpen] = useState(false);
    const pathname = usePathname();
    const hasSubpaths = route.subpaths && route.subpaths.length > 0;


    const hasSubItemToList = (route: IRoute) => {
        return route.subpaths?.some(path =>
            path.allowedRoles.some(allowedRoleItem => roleDictionary[allowedRoleItem] === user?.roleId) &&
            path.renderNavigator
        );
    };

    const hasAccesToItem = () => {
        return route.allowedRoles.some((role) => {
            return roleDictionary[role] === user?.roleId
        })
    }

    const isActive = pathname.includes(`/manager/${route.path}`);

    return (
       <>
       {
        hasAccesToItem() ?  <div className="mb-2 z-40">
            {route.renderNavigator && (
                <>
                    {hasSubpaths && hasSubItemToList(route) ? (
                        <>
                            <div
                                className={`flex justify-between items-center btn-primary ${isActive ? "border-1 border-white" : ""}`}
                                onClick={() => setOpen(!open)}
                            >
                                <div className="flex items-center gap-2">
                                    {React.cloneElement(route.icon, { size: 20 })}
                                    <span>{route.label}</span>
                                </div>
                                <IoMdArrowDropup
                                    size={20}
                                    className={`transition-transform duration-300 ${open ? "rotate-180" : "rotate-0"}`}
                                />
                            </div>

                            <div
                                className="overflow-hidden transition-all duration-300 ease-in-out"
                                style={{
                                    maxHeight: open ? `${route.subpaths!.length * 3}rem` : "0",
                                }}
                            >
                                {route.subpaths!.map((sub, subIndex) => (
                                    <SidebarSubItem
                                        key={subIndex}
                                        parentPath={route.path}
                                        subpath={sub}
                                    />
                                ))}
                            </div>
                        </>
                    ) : (
                        <Link
                            href={`/manager/${route.path}`}
                            className={`flex justify-between btn-primary ${isActive ? "border-1 border-white" : ""}`}
                        >
                            <div className="flex items-center gap-2">
                                {React.cloneElement(route.icon, { size: 20 })}
                                <span>{route.label}</span>
                            </div>
                        </Link>
                    )}
                </>
            )}
        </div>:null
       }
       </>
    );
}

interface SidebarSubItemProps {
    parentPath: string;
    subpath: IRoute;
}

const hasAccesToSubItem = (subpath: IRoute) => {
    const { user } = useAuth();
    return subpath.allowedRoles.some(role => roleDictionary[role] === user?.roleId);
};

const SidebarSubItem: React.FC<SidebarSubItemProps> = ({ parentPath, subpath }) => {
    const pathname = usePathname();
    const fullPath = `/manager/${parentPath}/${subpath.path}`;
    const isActive = pathname === fullPath;

    return (hasAccesToSubItem(subpath) && subpath.renderNavigator) ? (
        <Link
            href={fullPath}
            className={`flex items-center gap-2 text-sm text-white rounded-sm m-1 p-2 
                ${isActive ? "bg-primary" : "bg-primary/50 hover:bg-white/50"}`}
        >
            {React.cloneElement(subpath.icon, { size: 20 })}
            {subpath.label}
        </Link>
    ) : null;
};
