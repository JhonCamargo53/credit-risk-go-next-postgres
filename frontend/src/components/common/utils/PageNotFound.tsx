import React from "react";
import Link from "next/link";
import { useAuth } from "@/hooks/useAuth";

const PageNotFound = () => {

  const { user } = useAuth()

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100 px-4">
      <div className="text-center">
        <h1 className="text-9xl font-extrabold text-gray-400">404</h1>
        <p className="text-2xl font-semibold text-gray-700 mt-4">
          Página no encontrada
        </p>
        <p className="text-gray-500 mt-2">
          Lo sentimos, la página que buscas no existe o fue movida.
        </p>
        <Link
          href={user ? '/manager/dashboard' : '/login'}
          className="primary-gradient-to-t mt-6 inline-block bg-teal-600 text-white px-6 py-3 rounded-lg shadow hover:bg-teal-700 transition"
        >
          Volver al inicio
        </Link>
      </div>
    </div>
  );
};

export default PageNotFound;