"use client";

interface RiskLevelPillProps {
    score: number;
};

function getRiskStyles(score: number) {

    if (score >= 75) {

        return {
            label: "Riesgo Bajo", classes: "bg-emerald-500/10 text-emerald-700 border border-emerald-500/30",
        };
    }

    if (score >= 55) {
        return {
            label: "Riesgo Medio", classes: "bg-amber-500/10 text-amber-700 border border-amber-500/30",
        };
    }

    return {
        label: "Riesgo Alto", classes: "bg-red-500/10 text-red-700 border border-red-500/30",
    };
}

export const RiskLevelPill: React.FC<RiskLevelPillProps> = ({ score }) => {

    const {classes } = getRiskStyles(score);

    return (
        <span
            className={`inline-flex items-center gap-1 rounded-full px-3 py-1 text-xs font-semibold ${classes}`}
        >
            <span className="text-[11px] opacity-80">
                ({score.toFixed(1)}/100)
            </span>
        </span>
    );
};
