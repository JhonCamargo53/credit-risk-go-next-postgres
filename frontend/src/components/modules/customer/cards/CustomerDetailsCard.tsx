import React from "react";
import { Customer } from "@/types/customer";
import { useDocumentTypeStore } from "@/store/useDocumentTypeStore";

interface CustomerDetailsCardProps {
  customer: Customer;
}

const CustomerDetailsCard: React.FC<CustomerDetailsCardProps> = ({ customer }) => {

  const { documentTypes } = useDocumentTypeStore();
  const documentType = documentTypes.find(dt => dt.ID === customer.documentTypeId);

  const renderField = (label: string, value?: string | number | null) => (
    <p className="text-gray-500">
      <span className="text-black font-bold">{label}:</span>{" "}
      {value ? value : "───────────"}
    </p>
  );

  return (
    <div className="p-6 space-y-8 border rounded-xl shadow-md bg-white/60">
      <h1 className="text-4xl font-bold text-primary border-b pb-2">
        Información del Cliente
      </h1>

      {/* Información Personal */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <h2 className="text-xl font-semibold text-primary mb-4">
            Datos Generales
          </h2>
          {renderField("Nombre", customer.name)}
          {renderField("Correo Electrónico", customer.email)}
        </div>

        <div>
          <h2 className="text-xl font-semibold text-primary mb-4">
            Identificación
          </h2>
          {renderField("Tipo de Documento", documentType ? `${documentType.code} - ${documentType.description}` : "")}
          {renderField("Número de Documento", customer.documentNumber)}
        </div>
      </div>

      {/* Contacto */}
      <div>
        <h2 className="text-xl font-semibold text-primary mb-4">
          Información de Contacto
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-2">
          {renderField("Número de Teléfono", customer.phoneNumber)}
        </div>
      </div>

      {/* Información Financiera */}
      <div>
        <h2 className="text-xl font-semibold text-primary mb-4">
          Información Financiera
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-2">
          {renderField(
            "Ingresos Mensuales",
            customer.monthlyIncome
              ? "$" + customer.monthlyIncome.toLocaleString()
              : undefined
          )}
         
        </div>
      </div>
    </div>
  );
};

export default CustomerDetailsCard;
