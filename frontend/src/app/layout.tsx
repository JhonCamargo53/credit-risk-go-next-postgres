import type { Metadata } from "next";
import "../styles/globals.css";
import { AuthProvider } from "@/context/AuthContext";
import { Toaster } from 'react-hot-toast';
import LayoutWithAuth from "./layout-with-auth";
import { ModalProvider } from "@/context/ModalContext";

export const metadata: Metadata = {
  title: "Credit Risk Manager – MVP",
  description:
    "MVP del sistema inteligente para gestión de solicitudes de crédito, evaluación automática de riesgo, generación de reportes financieros y explicación en lenguaje natural.",
  keywords: [
    "credit risk",
    "evaluación de riesgo",
    "MVP",
    "gestor de crédito",
    "finanzas",
    "score crediticio",
    "solicitudes de crédito",
    "análisis financiero",
    "reportes automáticos",
    "machine learning",
    "Golang backend",
    "Next.js frontend",
    "PostgreSQL",
  ],
  authors: [
    { name: "Jhon Camargo" },
  ],
  applicationName: "Credit Risk Manager",
  category: "Finanzas / Análisis de Riesgo",
  generator: "Next.js",
  creator: "Jhon Camargo",
  publisher: " Prueba Tecnica -Ingeniería de Sistemas",
  robots: "index, follow",
  openGraph: {
    title: "Credit Risk Manager – MVP",
    description:
      "Sistema MVP para evaluar riesgo crediticio, gestionar solicitudes y generar reportes financieros automatizados.",
    url: "https://risk-management.alphacodexs.com",
    siteName: "Credit Risk Manager",
    locale: "es_CO",
    type: "website",
  }
};

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
      <body>
        <AuthProvider>
          <ModalProvider>
            <LayoutWithAuth>
              {children}
            </LayoutWithAuth>
            <Toaster />
          </ModalProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
