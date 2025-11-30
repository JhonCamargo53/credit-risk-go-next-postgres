import Button from '@/components/common/buttons/Button';
import GenericInput from '@/components/common/inputs/GenericInput';
import { generateAxiosErrorToast, showSuccessToast } from '@/utils/toastUtils';
import React, { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useModal } from '@/hooks/useModal';
import { useCreditRequest } from '@/hooks/useCreditRequest';
import { useCreditStatusStore } from '@/store/useCreditStatusStore';

interface CreateUpdateCreditRequestProps {
  creditRequest?: CreditRequest;
  customerId?: number;
}

const CreateUpdateCreditRequestForm: React.FC<CreateUpdateCreditRequestProps> = ({ creditRequest, customerId }) => {

  const { creditStatuses } = useCreditStatusStore();
  const { createCreditRequest, updateCreditRequest } = useCreditRequest();
  const { removeLastOpenModal } = useModal();
  const [loading, setLoading] = useState(false);

  const { register, handleSubmit, formState: { errors } } = useForm<CreditRequestForm>({
    defaultValues: {
      amount: creditRequest?.amount,
      termMonths: creditRequest?.termMonths,
      productType: creditRequest?.productType,
      creditStatusId: creditRequest?.creditStatusId,
    }
  });

  const onSubmit: SubmitHandler<CreditRequestForm> = async (formData) => {
    try {
      setLoading(true);

      const payload: CreditRequestForm = {
        ...formData,
        amount: Number(formData.amount),
        termMonths: Number(formData.termMonths),
        creditStatusId: Number(formData.creditStatusId),
        customerId: customerId ? customerId : (creditRequest ? creditRequest.ID : 0),
      }

      if (creditRequest) {
        await updateCreditRequest(creditRequest.ID, payload);
      } else {
        await createCreditRequest(payload);
      }

      removeLastOpenModal();
      showSuccessToast('Solicitud guardada con éxito', `La solicitud ha sido ${creditRequest ? 'actualizada' : 'creada'} correctamente.`);

    } catch (error) {
      generateAxiosErrorToast(error, 'Error al guardar solicitud', 'Inténtalo nuevamente');
    } finally {
      setLoading(false);
    }
  };

  return (
    <form className="flex flex-col justify-between h-[calc(100%-32px)]" onSubmit={handleSubmit(onSubmit)}>
      <div className="flex-1 space-y-8">
        <div className="space-y-4">
          <div className="grid grid-cols-1 gap-4">
            <GenericInput
              label="Monto"
              placeholder="Ejemplo: 500000"
              type="number"
              error={errors.amount}
              register={register('amount', {
                required: 'El monto es obligatorio',
                min: { value: 0, message: 'El monto no puede ser negativo' }
              })}
            />
          </div>

          <div className="grid grid-cols-1 gap-4">
            <GenericInput
              label="Plazo (meses)"
              placeholder="Ejemplo: 12"
              type="number"
              error={errors.termMonths}
              register={register('termMonths', {
                required: 'El plazo es obligatorio',
                min: { value: 1, message: 'El plazo debe ser al menos de 1 mes' }
              })}
            />
          </div>

          <div className="grid grid-cols-1 gap-4">
            <GenericInput
              label="Tipo de producto"
              placeholder="Ejemplo: Crédito personal"
              type="text"
              error={errors.productType}
              register={register('productType', {
                required: 'El tipo de producto es obligatorio',
                maxLength: { value: 50, message: 'No puede tener más de 50 caracteres' }
              })}
            />
          </div>

          <div>
            <label className="text-neutral-dark font-bold mb-1 block">Estado del crédito</label>
            <select
              className="input-primary w-full h-10"
              {...register('creditStatusId', {
                required: 'El estado es obligatorio'
              })}
            >
              <option value="">Selecciona un estado</option>
              {creditStatuses.map(status => (
                <option key={status.ID} value={status.ID}>
                  {status.name}
                </option>
              ))}
            </select>
            {errors.creditStatusId && (
              <p className="text-sm text-error font-bold">{errors.creditStatusId.message}</p>
            )}
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

export default CreateUpdateCreditRequestForm;
