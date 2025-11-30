'use client';
import { IoLogOut, IoMenu } from "react-icons/io5";
import { useState } from "react";
import { useAuth } from "@/hooks/useAuth";
import PageNotFound from "@/components/common/utils/PageNotFound";
import Sidebar from "@/components/common/sidebar/Sidebar";
import LoadingPage from "@/components/common/loading/LoadingPage";

export default function ManagerLayout({ children }: Readonly<{ children: React.ReactNode; }>) {
  const [open, setOpen] = useState(true);
  const { user, logout, loadingLogout } = useAuth()

  if (loadingLogout) return <LoadingPage loadingText='Cerrando sesión' />

  if (!user) return <PageNotFound />

  return (
    <div className="flex flex-col min-h-screen h-full">
      <div className="bg-neutral-dark h-[60px] flex items-center justify-between  px-4  w-full " />
      <div className="bg-neutral-dark h-[60px] flex items-center justify-between  px-4  w-full fixed z-40">
        <IoMenu onClick={() => setOpen(!open)} size={30} className="cursor-pointer text-white" />

        <button title="Cerrar Sesión"><IoLogOut onClick={() => logout()} size={25} className="cursor-pointer text-white hover:text-gray-300 transition-colors duration-200" /></button>
      </div>
      <div className="flex flex-1 h-full pt-0"
        style={{
          marginLeft: open ? '16rem' : '0',
          transition: 'margin-left 0.3s ease-in-out'
        }}>

        <Sidebar open={open} />

        <div className={`flex-1 bg-neutral-dark pr-3 pb-3 ${!open ? 'pl-3' : 'pl-0'}`}>
          <div className="bg-neutral-light rounded p-4 h-full">
            {children}
          </div>
        </div>
      </div>
    </div>
  );
}
