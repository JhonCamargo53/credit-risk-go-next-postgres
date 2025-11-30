import Button from '@/components/common/buttons/Button';
import GenericInput from '@/components/common/inputs/GenericInput';
import { generateAxiosErrorToast, showSuccessToast } from '@/utils/toastUtils';
import React, { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useModal } from '@/hooks/useModal';
import { Customer, CustomerForm } from '@/types/customer';
import { useDocumentTypeStore } from '@/store/useDocumentTypeStore';
import { useCustomer } from '@/hooks/useCustomer';

interface CreateUpdateCustomerFormProps {
  customer?: Customer
}

const CreateUpdateCustomerForm: React.FC<CreateUpdateCustomerFormProps> = ({ customer }) => {

  const { documentTypes } = useDocumentTypeStore();
  const { createCustomer, updateCustomer } = useCustomer();
  const { removeLastOpenModal } = useModal();
  const [loading, setLoading] = useState(false);

  const { register, handleSubmit, formState: { errors }, watch } = useForm<CustomerForm>({
    defaultValues: {
      name: customer?.name,
      email: customer?.email,
      phoneNumber: customer?.phoneNumber,
      documentNumber: customer?.documentNumber,
      documentTypeId: customer?.documentTypeId,
      monthlyIncome: customer?.monthlyIncome,
    }
  });

  const onSubmit: SubmitHandler<CustomerForm> = async (formData) => {
    try {

      setLoading(true);

      const payload = { ...formData, documentTypeId: Number(formData.documentTypeId), monthlyIncome: Number(formData.monthlyIncome) };
      customer ? await updateCustomer(customer.ID, payload) : await createCustomer(payload);

      removeLastOpenModal()
      showSuccessToast('Usuario guardado con éxito', `El cliente ha sido ${customer ? 'actualizado' : 'creado'} correctamente.`);

    } catch (error) {
      generateAxiosErrorToast(error, 'Error al guardar cliente', 'Intentelo nuevamente');
    } finally {
      setLoading(false);
    }
  };

  return (
    <form className="flex flex-col justify-between h-[calc(100%-32px)]" onSubmit={handleSubmit(onSubmit)}    >
      <div className="flex-1 space-y-8">

        <div className='space-y-4'>
          <div className="grid grid-cols-1 gap-4">
            <GenericInput
              label="Nombre"
              placeholder="Ejemplo: Juan Ahumada"
              type="text"
              error={errors.name}
              register={register('name', {
                required: 'El nombre es obligatorio',
                maxLength: { value: 54, message: 'No puede tener más de 54 caracteres.' }
              })}
            />
          </div>

          <div className="grid grid-cols-1  gap-4">
            <GenericInput
              label="Correo electrónico"
              placeholder="Ejemplo: management@gmail.com"
              type="email"
              error={errors.email}
              register={register('email', {
                required: 'El correo electrónico es obligatorio.',
                pattern: { value: /\S+@\S+\.\S+/, message: 'Correo electrónico no válido.' }
              })}
            />
          </div>

          <div className="grid grid-cols-1  gap-4">
            <GenericInput
              label="Número de teléfono"
              placeholder="Ejemplo: 301545*****"
              type="tet"
              error={errors.phoneNumber}
              register={register('phoneNumber', {
                required: 'Número de teléfono es obligatorio.',
                maxLength: { value: 15, message: 'No puede tener más de 15 caracteres.' }
              })}
            />
          </div>

          <div className="grid grid-cols-1 gap-4">
            <GenericInput
              label="Número de documento"
              placeholder="Ejemplo: 1332134****"
              type="tet"
              error={errors.phoneNumber}
              register={register('documentNumber', {
                required: 'Número de documento es obligatorio.',
                maxLength: { value: 20, message: 'No puede tener más de 20 caracteres.' }
              })}
            />
          </div>

          <div>
            <label className='text-neutral-dark font-bold mb-1 block'>Tipo de Documento de Identidad</label>
            <select
              className="input-primary w-full h-10"
              {...register('documentTypeId', {
                required: 'El rango es obligatorio'
              })}
            >
              <option value="">Selecciona un tipo de documento</option>
              {documentTypes.map(documentType => (
                <option key={documentType.ID} value={documentType.ID}>
                  {documentType.code} - {documentType.description}
                </option>
              ))}
            </select>
            {errors.documentNumber && (
              <p className="text-sm text-error font-bold">{errors.documentNumber.message}</p>
            )}
          </div>


          <div className="grid grid-cols-1 gap-4">
            <GenericInput
              label="Ganancia mensual"
              placeholder="Ejemplo: 1332134****"
              type="number"
              error={errors.monthlyIncome}
              register={register('monthlyIncome', {
                required: 'Ganancia mensual es obligatoria.',
                min: { value: 0, message: 'La ganancia mensual no puede ser negativa.' },
                maxLength: { value: 15, message: 'No puede tener más de 15 caracteres.' }
              })}
            />
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

export default CreateUpdateCustomerForm;
