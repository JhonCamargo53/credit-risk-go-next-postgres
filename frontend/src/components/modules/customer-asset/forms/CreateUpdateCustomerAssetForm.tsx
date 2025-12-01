import Button from '@/components/common/buttons/Button';
import GenericInput from '@/components/common/inputs/GenericInput';
import { useCreditRequest } from '@/hooks/useCreditRequest';
import { useCustomerAsset } from '@/hooks/useCustomerAsset';
import { useModal } from '@/hooks/useModal';
import { useAssetStore } from '@/store/useAssetStore';
import { CustomerAsset, CustomerAssetForm } from '@/types/customerAsset';
import { generateAxiosErrorToast, showSuccessToast } from '@/utils/toastUtils';
import React, { useState } from 'react'
import { SubmitHandler, useForm } from 'react-hook-form';


interface CreateUpdateCustomerAssetFormProps {
  customerAsset?: CustomerAsset;
  customerId: number;
  creditRequestId: number;
}

const CreateUpdateCustomerAssetForm: React.FC<CreateUpdateCustomerAssetFormProps> = ({
  customerAsset, customerId, creditRequestId }) => {

  const { assets } = useAssetStore()
  const { fetchCreditRequestById, addOrUpdateCreditRequest } = useCreditRequest()
  const { createCustomerAsset, updateCustomerAsset } = useCustomerAsset();
  const { removeLastOpenModal } = useModal();
  const [loading, setLoading] = useState(false);

  const { register, handleSubmit, formState: { errors } } = useForm<CustomerAssetForm>({
    defaultValues: {
      description: customerAsset?.description,
      marketValue: customerAsset?.marketValue,
      assetId: customerAsset?.assetId,
      creditRequestId: creditRequestId || customerAsset?.creditRequestId,
      customerId: customerId || customerAsset?.customerId,
    }
  });

  const onSubmit: SubmitHandler<CustomerAssetForm> = async (formData) => {
    try {
      setLoading(true);

      const payload: CustomerAssetForm = {
        ...formData,
        marketValue: Number(formData.marketValue),
        creditRequestId: Number(formData.creditRequestId),
        assetId: Number(formData.assetId),
        customerId: Number(formData.customerId),

      };


      if (customerAsset) {
        await updateCustomerAsset(customerAsset.ID, payload);
        showSuccessToast('Activo actualizado', 'El activo del cliente ha sido actualizado correctamente.');
      } else {
        await createCustomerAsset(payload);
        showSuccessToast('Activo creado', 'El activo del cliente ha sido creado correctamente.');
      }

      const creditRequest = await fetchCreditRequestById(creditRequestId)
      addOrUpdateCreditRequest(creditRequest);
      removeLastOpenModal();
    } catch (error) {
      generateAxiosErrorToast(error, 'Error al guardar activo', 'Inténtalo nuevamente');
    } finally {
      setLoading(false);
    }
  };

  return (
    <form className="flex flex-col justify-between h-[calc(100%-32px)]" onSubmit={handleSubmit(onSubmit)}>
      <div className="flex-1 space-y-8">
        <div className="space-y-4">
          <GenericInput
            label="Descripción"
            placeholder="Ejemplo: Auto Toyota Corolla"
            type="text"
            error={errors.description}
            register={register('description', {
              required: 'La descripción es obligatoria',
              maxLength: { value: 100, message: 'No puede tener más de 100 caracteres' },
            })}
          />

          <GenericInput
            label="Valor de mercado"
            placeholder="Ejemplo: 25000000"
            type="number"
            error={errors.marketValue}
            register={register('marketValue', {
              required: 'El valor de mercado es obligatorio',
              min: { value: 0, message: 'El valor no puede ser negativo' },
            })}
          />

          <div>
            <label className="text-neutral-dark font-bold mb-1 block">Tipo de Activo</label>
            <select
              className="input-primary w-full h-10"
              {...register('assetId', {
                required: 'El tipo de activo es obligatorio',
              })}
              defaultValue={customerAsset?.assetId || ""}
            >
              <option value="">Selecciona un tipo de activo</option>
              {assets.map((asset) => (
                <option key={asset.ID} value={asset.ID}>
                  {asset.name}
                </option>
              ))}
            </select>
            {errors.assetId && (
              <p className="text-sm text-error font-bold">{errors.assetId.message}</p>
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


export default CreateUpdateCustomerAssetForm