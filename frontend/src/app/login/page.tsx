'use client'
import { useState } from "react";
import { SubmitHandler, useForm } from 'react-hook-form';
import { useAuth } from "@/hooks/useAuth";
import { useRouter } from "next/navigation";
import { generateAxiosErrorToast } from "@/utils/toastUtils";
import GenericInput from "@/components/common/inputs/GenericInput";
import Button from "@/components/common/buttons/Button";
import { LoginForm } from "@/types/auth";

export default function LoginPage() {

  const { login } = useAuth();
  const { register: registerLogin, handleSubmit, formState: { errors } } = useForm<LoginForm>();
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const onSubmit: SubmitHandler<LoginForm> = async (formData: LoginForm) => {
    try {
      setLoading(true);
      await login(formData);
      router.replace("/manager/dashboard");
    } catch (error: unknown) {
      generateAxiosErrorToast(error, 'Error al iniciar sesión', 'Intentelo nuevamente');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="m-4 grid grid-cols-1 lg:grid-cols-9 min-h-[90vh] lg:min-h-[70vh] lg:shadow-xl max-w-[90vw] lg:max-w-[65vw] rounded-4xl overflow-hidden">

      <div className="bg-neutral-light/90 flex flex-col justify-center items-center p-10 col-span-4">
        <div className="w-full max-w-md">
          <h3 className="font-bold text-4xl text-center text-neutral-dark">
            Accede al Sistema de Créditos
          </h3>

          <p className="text-neutral-dark/70 text-center mt-2">
            Inicia sesión para gestionar clientes, solicitudes de crédito y evaluaciones de riesgo.
          </p>

          <form className="flex flex-col gap-4 mt-8" onSubmit={handleSubmit(onSubmit)}>
            <GenericInput
              placeholder="Correo electrónico"
              type="text"
              error={errors.email}
              register={registerLogin('email', {
                required: 'El correo es obligatorio',
                pattern: {
                  value: /^\S+@\S+\.\S+$/,
                  message: 'Correo no válido',
                },
              })}
            />

            <GenericInput
              placeholder="Contraseña"
              type="password"
              error={errors.password}
              register={{
                ...registerLogin('password', {
                  required: 'La contraseña es obligatoria',
                })
              }}
            />

            <div className="flex justify-center mt-6">
              <Button
                type="submit"
                loading={loading}
                className="w-full lg:w-64 from-primary to-primary-dark hover:scale-105 transition-transform duration-300"
              >
                Iniciar sesión
              </Button>
            </div>
          </form>
        </div>
      </div>

      <div className="bg-radial-primary-to-dark flex flex-col justify-center items-center col-span-5 px-12 py-10 text-white">
        <h2 className="text-4xl font-extrabold mb-4 text-center">
          Sistema de Gestión
        </h2>

        <p className="text-lg leading-relaxed text-justify">
          Administra clientes, registra solicitudes de crédito y realiza
          Toma decisiones informadas y optimiza el proceso de evaluación crediticia.
        </p>
      </div>
    </div>
  );
}
