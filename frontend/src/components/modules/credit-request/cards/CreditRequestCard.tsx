import { useCreditStatusStore } from "@/store/useCreditStatusStore";
import DOMPurify from "dompurify";
import React from "react";

interface CreditRequestCardProps {
    creditRequest: CreditRequest;
}

const CreditRequestCard: React.FC<CreditRequestCardProps> = ({ creditRequest }) => {

    const { creditStatuses } = useCreditStatusStore();

    const creditStatus = creditStatuses.find(status => status.ID === creditRequest.creditStatusId);

    const renderField = (label: string, value?: string | number | null) => (
        <p className="text-gray-500">
            <span className="text-black font-bold">{label}:</span>{" "}
            {value !== null && value !== undefined ? value : "───────────"}
        </p>
    );

    return (
        <div className="p-6 space-y-8 border rounded-xl shadow-md bg-white/60">
            <h1 className="text-4xl font-bold text-primary border-b pb-2">
                Solicitud de Crédito
            </h1>

            {/* Fechas */}
            <div className="grid grid-cols-1 md:grid-cols-2 ">
                <div>
                    <h2 className="text-xl font-semibold text-primary mb-4">
                        Fechas
                    </h2>
                    <div className="col-span-6">
                        {renderField("Creación", new Date(creditRequest.CreatedAt).toLocaleString())}
                    </div>
                    <div>
                        {renderField("Última Actualización", new Date(creditRequest.UpdatedAt).toLocaleString())}

                    </div>
                </div>
                <div>
                    <div>
                        <h2 className="text-xl font-semibold text-primary mb-4">
                            Detalles Generales
                        </h2>
                        {renderField("Monto", `$${creditRequest.amount.toLocaleString()}`)}
                        {renderField("Plazo (meses)", creditRequest.termMonths)}
                        {renderField("Tipo de Producto", creditRequest.productType)}
                        {renderField("Estado del Crédito", creditStatus?.name)}
                    </div>
                </div>
            </div>


            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

                <div className="col-span-12">
                    <h2 className="text-xl font-semibold text-primary mb-4">
                        Evaluación de Riesgo
                    </h2>
                  

                    <div>
                        {creditRequest.riskExplanation.split("-").slice(1).map((paragraph, index) => {
                            const parts = paragraph.trim().split(/(\{[^}]*\})/g);
                            return (
                                <p key={index} className="text-gray-600 mb-2 text-justify">
                                    <span className="text-primary font-bold">{index+1}.</span>{" "}
                                    {parts.map((part, i) => 
                                        part.startsWith("{") && part.endsWith("}") ? (
                                            <span key={i} className="text-black font-bold">{part.replace("{","").replace("}","")}</span>
                                        ) : (
                                            <span key={i}>{part}</span>
                                        )
                                    )}
                                </p>
                            );
                        })}
                    </div>
                </div>
            </div>
        </div>
    );
};


interface RiskExplanationProps {
    riskExplanation: string;
}

export const RiskExplanation: React.FC<RiskExplanationProps> = ({ riskExplanation }) => {

    const html = riskExplanation ?? "";

    const createMarkup = (dirty: string) => ({
        __html: DOMPurify.sanitize(dirty),
    });

    return (
        <div>
            <h3>Explicación de riesgo</h3>
            <div
                className="risk-explanation"
                dangerouslySetInnerHTML={createMarkup(html)}
            />
        </div>
    );
}


export default CreditRequestCard;
