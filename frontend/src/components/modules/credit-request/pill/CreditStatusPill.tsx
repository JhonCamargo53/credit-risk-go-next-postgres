"use client";

interface CreditStatusPillProps {
    statusId: number;
};

function getStatusById(statusId: number) {
    switch (statusId) {
        case 2:
            return {
                label: "Aprobado",
                classes:
                    "bg-emerald-500/10 text-emerald-700 border border-emerald-500/30",
            };
        case 3:
            return {
                label: "Rechazado",
                classes:
                    "bg-red-500/10 text-red-700 border border-red-500/30",
            };
        case 4:
            return {
                label: "En estudio",
                classes:
                    "bg-sky-500/10 text-sky-700 border border-sky-500/30",
            };
        case 1:
        default:
            return {
                label: "Pendiente",
                classes:
                    "bg-slate-500/10 text-slate-700 border border-slate-500/30",
            };
    }
}

export const CreditStatusPill: React.FC<CreditStatusPillProps> = ({ statusId, }) => {
  
    const { label, classes } = getStatusById(statusId);

    return (
        <span className={`inline-flex items-center rounded-full px-3 py-1 text-xs font-semibold ${classes}`}        >
            {label}
        </span>
    );
};
