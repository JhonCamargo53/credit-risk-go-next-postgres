import Button from '@/components/common/buttons/Button';
import GenericInput from '@/components/common/inputs/GenericInput';
import { generateAxiosErrorToast, showSuccessToast } from '@/utils/toastUtils';
import React, { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { User, UserForm } from '@/types/user';
import { useModal } from '@/hooks/useModal';
import { useUser } from '@/hooks/useUserContext';

interface CreateUpdateUserFormProps {
  user?: User
}

const CreateUpdateUserForm: React.FC<CreateUpdateUserFormProps> = ({ user }) => {

  const { createUser, updateUser } = useUser();
  const { removeLastOpenModal } = useModal();
  const [loading, setLoading] = useState(false);
  const [roles] = useState([{ id: 1, name: 'ADMINISTRADOR' }, { id: 2, name: 'EMPLEADO' }])

  const { register, handleSubmit, formState: { errors },watch } = useForm<UserForm>({
    defaultValues: {
      name: user?.name,
      roleId: user?.roleId,
      email: user?.email,
      password: ''
    }
  });

  const onSubmit: SubmitHandler<UserForm> = async (formData) => {
    try {

      setLoading(true);

      user
        ? await updateUser(user.ID, { ...formData, roleId: Number(formData.roleId) })
        : await createUser({ ...formData, roleId: Number(formData.roleId) });

      removeLastOpenModal();
      showSuccessToast('Usuario guardado con éxito', `El usuario ha sido ${user ? 'actualizado' : 'creado'} correctamente.`);

    } catch (error) {
      generateAxiosErrorToast(error, 'Error al guardar usuario', 'Intentelo nuevamente');
    } finally {
      setLoading(false);
    }
  };

  return (
    <form className="flex flex-col justify-between h-[calc(100%-32px)]" onSubmit={handleSubmit(onSubmit)}    >
      <div className="flex-1 space-y-8">
        {/* Información Personal */}
        <div className='space-y-4'>
          <div className="grid grid-cols-1 gap-4">
            <GenericInput
              label="Primer Nombre"
              placeholder="Ejemplo: Juan Ahumada"
              type="text"
              error={errors.name}
              register={register('name', {
                required: 'El nombre es obligatorio',
                maxLength: { value: 54, message: 'No puede tener más de 54 caracteres.' }
              })}
            />
          </div>

          {/* Información de Ingreso */}
          <div>
            <div className="grid grid-cols-1  gap-4">
              <GenericInput
                label="Correo electrónico"
                placeholder="management@gmail.com"
                type="email"
                error={errors.email}
                register={register('email', {
                  required: 'El correo electrónico es obligatorio.',
                  pattern: { value: /\S+@\S+\.\S+/, message: 'Correo electrónico no válido.' }
                })}
              />
            </div>
          </div>

          <div className='space-y-4'>
            <div className="grid grid-cols-1 gap-4">
             <div>
                <label className='text-neutral-dark font-bold mb-1 block'>Rango</label>
                <select
                  className="input-primary w-full h-10"
                  {...register('roleId', {
                    required: 'El rango es obligatorio'
                  })}
                >
                  <option value="">Selecciona un rango</option>
                  {roles.map(role => (
                    <option key={role.id} value={role.id}>
                      {role.name}
                    </option>
                  ))}
                </select>
                {errors.roleId && (
                  <p className="text-sm text-error font-bold">{errors.roleId.message}</p>
                )}
              </div>
              <div>
                <GenericInput
                  label={'Contraseña'}
                  placeholder="Contraseña"
                  type="password"
                  error={errors.password}
                  register={register('password', {
                    required: !user ? 'La contraseña es obligatoria' : false,
                  })}
                />
                {user ? <label className='italic text-gray-700'> (Dejar en blanco para mantener Actual)</label> : null}
              </div>

            </div>
          </div>
        </div>


        {/* Botones */}
        <div className="flex flex-col md:flex-row justify-end mt-6 gap-3 pb-2">
          <Button
            type="button"
            disabled={loading}
            variant="secondary"
            className="w-full md:w-40"
            onClick={() => removeLastOpenModal()}
          >
            Cancelar
          </Button>
          <Button type="submit" loading={loading} className="w-full md:w-40">
            Guardar
          </Button>
        </div>
      </div>
    </form>

  );
};

export default CreateUpdateUserForm;
