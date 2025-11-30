'use client'

import ProtectedRoute from "@/hoc/ProtectedRoute";
import { FaShieldAlt, FaUserCheck, FaChartLine, FaDatabase, FaServer, FaBrain } from "react-icons/fa";

const HomePage = () => {
     return (
          <div className="w-full h-[90vh]  text-neutral-light flex flex-col items-center justify-between px-8 py-6 rounded-md">

               <header className="w-full max-w-5xl flex items-center justify-center mt-20">
                    <h2 className="flex items-center gap-3 text-2xl font-extrabold tracking-widest text-primary drop-shadow-[0_0_8px_rgba(197,52,52,0.6)]">
                         <FaShieldAlt className="text-primary" size={28} />
                         JHON CAMARGO · DESARROLLADOR FULLSTACK
                         <FaBrain className="text-primary" size={28} />
                    </h2>
               </header>

               <main className="w-full flex items-center justify-center">
                    <div className="relative max-w-3xl w-full">

                         <div className="absolute inset-0 blur-3xl bg-primary/40 rounded-[2.5rem] opacity-60 -z-10" />

                         <div className="bg-linear-primary-to-dark rounded-[2.5rem] p-[1px] shadow-2xl">
                              <div className="bg-neutral-dark/60 backdrop-blur-xl border border-primary/30 rounded-[2.4rem] p-10 flex flex-col items-center">

                                   <div className="bg-primary/20 p-6 rounded-full border border-primary/40 shadow-lg">
                                        <FaShieldAlt className="text-primary" size={48} />
                                   </div>

                                   <h1 className="text-4xl font-bold mt-6 text-primary text-center">
                                        Gestor de Solicitudes de Crédito
                                   </h1>

                                   <p className="text-neutral-light/70 mt-3 text-center max-w-xl">
                                        Administra clientes, solicitudes y realiza pre-evaluaciones automáticas de riesgo mediante IA.
                                   </p>

                                   <div className="grid grid-cols-1 sm:grid-cols-3 gap-6 mt-10 w-full">

                                        <div className="bg-neutral-light/5 rounded-xl p-5 flex flex-col items-center border border-neutral-light/10 hover:bg-neutral-light/10 transition">
                                             <FaUserCheck className="text-secondary mb-2" size={40} />
                                             <p className="text-center text-sm text-neutral-light/80">
                                                  Clientes &amp; Solicitudes
                                             </p>
                                        </div>

                                        <div className="bg-neutral-light/5 rounded-xl p-5 flex flex-col items-center border border-neutral-light/10 hover:bg-neutral-light/10 transition">
                                             <FaChartLine className="text-accent mb-2" size={40} />
                                             <p className="text-center text-sm text-neutral-light/80">
                                                  Evaluación Inteligente
                                             </p>
                                        </div>

                                        <div className="bg-neutral-light/5 rounded-xl p-5 flex flex-col items-center border border-neutral-light/10 hover:bg-neutral-light/10 transition">
                                             <FaShieldAlt className="text-info mb-2" size={40} />
                                             <p className="text-center text-sm text-neutral-light/80">
                                                  Decisiones fiables
                                             </p>
                                        </div>

                                   </div>
                              </div>
                         </div>
                    </div>
               </main>

               <footer className="w-full max-w-5xl flex flex-wrap gap-3 justify-between text-xs text-primary font-bold mt-4">
                    <div className="flex items-center gap-2">
                         <FaServer /> <span>Backend: Go </span>
                    </div>
                    <div className="flex items-center gap-2">
                         <FaDatabase /> <span>DB: PostgreSQL</span>
                    </div>
                    <div className="flex items-center gap-2">
                         <FaBrain /> <span>Motor de riesgo crediticio</span>
                    </div>
               </footer>
          </div>
     );
};

export default ProtectedRoute(HomePage);
