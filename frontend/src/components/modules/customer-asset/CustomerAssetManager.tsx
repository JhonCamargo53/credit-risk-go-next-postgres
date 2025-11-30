import Button from '@/components/common/buttons/Button';
import GenericModal from '@/components/common/modals/GenericModal';
import GenericTable from '@/components/common/tables/GenericTable'
import { DataColumn } from '@/types/table';
import React, { useEffect, useState } from 'react'
import { FaPlus } from 'react-icons/fa';
import { useCustomerAsset } from '@/hooks/useCustomerAsset';
import CustomerAssetRow from './rows/CustomerAssetRow';
import CreateUpdateCustomerAssetForm from './forms/CreateUpdateCustomerAssetForm';
import { useAssetStore } from '@/store/useAssetStore';
import { useCreditRequest } from '@/hooks/useCreditRequest';
import { generateAxiosErrorToast } from '@/utils/toastUtils';

interface CustomerAssetManagerProps {
  creditRequestId: number
}

const userColumns: readonly DataColumn[] = [
  { id: 'index', label: 'N°' },
  { id: 'description', label: 'Descripción' },
  { id: 'marketValue', label: 'Valor de mercado', minWidth: 150 },
  { id: 'assetType', label: 'Tipo de Activo', minWidth: 150 },
  { id: 'edit', label: 'Editar' },
  { id: 'delete', label: 'Eliminar' }
];

const CustomerAssetManager: React.FC<CustomerAssetManagerProps> = ({ creditRequestId }) => {

  const [createUpdateModal, setCreateUpdateModal] = useState(false);
  const { creditRequests } = useCreditRequest()
  const { fetchAssets } = useAssetStore()
  const { customerAssets, loading, pageData, handlePagination, fetchCustomerAssetsByCreditRequestId } = useCustomerAsset();

  useEffect(() => {
    loadCustomerAssets()
  }, [])

  const loadCustomerAssets = async () => {
    try {
      await fetchAssets()
      await fetchCustomerAssetsByCreditRequestId(creditRequestId)
    } catch (error) {
      generateAxiosErrorToast(error, 'Error al cargar activos del usuario', 'Inténtalo nuevamente');
    }
  }

  const customerId = creditRequests.find(cr => cr.ID === creditRequestId)?.customerId!

  return (
    <div className="space-y-4">
      <div className="grid grid-cols-12 gap-4">
        <div className="col-span-2">
          <Button iconLeft={<FaPlus />} className="font-bold w-full h-full" onClick={() => setCreateUpdateModal(true)}>Agregar</Button>
        </div>
      </div>

      <div>
        <GenericTable
          columns={userColumns}
          data={customerAssets}
          row={<CustomerAssetRow />}
          handlePagination={handlePagination}
          loading={loading.fetching}
          pageData={pageData}
        />
      </div>

      {createUpdateModal && <GenericModal content={<CreateUpdateCustomerAssetForm customerId={customerId} creditRequestId={creditRequestId} />} isOpen={createUpdateModal} setOpen={setCreateUpdateModal} size="2xl" />}

    </div>

  )
}

export default CustomerAssetManager