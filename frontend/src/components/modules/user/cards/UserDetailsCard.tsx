import { User } from '@/types/user'
import { formatDateToDDMMYYYY } from '@/utils/dateUtils';
import React from 'react'
import { BsPersonBadge } from "react-icons/bs";

interface UserDetailsCardProps {
    user: User
}

const UserDetailsCard: React.FC<UserDetailsCardProps> = ({ user }) => {

    const renderField = (label: string, value?: string) => (
        <p className='text-gray-500'>
            <span className=" text-black font-bold">{label}:</span>{" "}
            {value?.trim() ? value : "───────────"}
        </p>
    );

    return (
        <div className="p-6 space-y-6 border rounded-xl shadow-md bg-white/60">

            <h1 className="text-2xl font-bold text-primary border-b pb-2">
                Información del Usuario
            </h1>

            {/* Tarjeta de Información Personal */}
            <div className="space-y-4">
                <h2 className="text-lg font-semibold text-primary">
                    Información Personal
                </h2>

                <div className="grid grid-cols-1  gap-4">
                    {renderField("Nombre completo", user?.name || "No disponible")}
                </div>
            </div>

            {/* Tarjeta de Contacto */}
            <div className="space-y-4">
                <h2 className="text-lg font-semibold text-primary">
                    Información de Contacto
                </h2>

                <div className="grid grid-cols-1  gap-4">
                    {renderField("Correo Electrónico", user?.email || "No disponible")}
                </div>
            </div>

        </div>
    );
}

export default UserDetailsCard