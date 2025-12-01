import Button from '@/components/common/buttons/Button';
import GenericModal from '@/components/common/modals/GenericModal';
import { GenericTableRow } from '@/types/table'
import React, { useState } from 'react'
import { FaPencilAlt } from 'react-icons/fa';
import { CustomerAsset } from '@/types/customerAsset';
import { useAssetStore } from '@/store/useAssetStore';
import CreateUpdateCustomerAssetForm from '../forms/CreateUpdateCustomerAssetForm';
import { RiDeleteBin6Fill } from "react-icons/ri";
import { generateAxiosErrorToast, showSuccessToast } from '@/utils/toastUtils';
import { useCustomerAsset } from '@/hooks/useCustomerAsset';
import { confirmActionAlert } from '@/utils/alertUtils';
import { useCreditRequest } from '@/hooks/useCreditRequest';

interface CustomerAssetRowProps extends GenericTableRow<CustomerAsset> {

}

const CustomerAssetRow: React.FC<CustomerAssetRowProps> = ({ index = 0, data, pageData = { currentPage: 1, limit: 10, lastPage: -1 } }) => {

    const [customerAsset] = useState(data!);

    const [customerAssetUpdateModal, setCustomerAssetUpdateModal] = useState(false);
    const { fetchCreditRequestById, addOrUpdateCreditRequest } = useCreditRequest()

    const { assets } = useAssetStore()

    const [deleting, setDeleting] = useState(false);
    const { deleteCustomerAsset } = useCustomerAsset()

    const asset = assets.find(asset => asset.ID === customerAsset.assetId);

    const handleDelete = async () => {
        try {

            const confirm = await confirmActionAlert('Eliminar activo', '¿Está seguro de realizar esta acción?', 'question')

            if (confirm) {
                setDeleting(true)
                await deleteCustomerAsset(customerAsset.ID);
                const creditRequest = await fetchCreditRequestById(customerAsset.creditRequestId)
                addOrUpdateCreditRequest(creditRequest);
                showSuccessToast('Eliminado', 'El activo del cliente ha sido eliminado correctamente.');
            }

        } catch (error) {
            generateAxiosErrorToast(error, 'Error al eliminar', 'Inténtalo nuevamente');

        } finally {
            setDeleting(false);
        }
    }

    return (
        <tr key={customerAsset.ID} className={`border-b border-gray-300 hover:bg-gray-200 text-center font-medium ${index % 2 !== 0 ? 'bg-primary/8' : ''}`}>
            <td className="px-4 py-2">{index + 1 + (pageData.currentPage - 1) * pageData.limit}</td>
            <td className="px-4 py-2 whitespace-nowrap">{customerAsset.description}</td>
            <td className="px-4 py-2 whitespace-nowrap">{customerAsset.marketValue.toLocaleString()}</td>
            <td className="px-4 py-2 whitespace-nowrap">{asset?.name}</td>
            <td className="px-4 py-2 ">
                <div className="flex items-center justify-center h-full">
                    <Button title="Editar activo" onClick={() => setCustomerAssetUpdateModal(true)}>
                        <FaPencilAlt />
                    </Button>
                </div>
            </td>
            <td className="px-4 py-2 ">
                <div className="flex items-center justify-center h-full">
                    <Button title="Eliminar activo" loading={deleting} onClick={() => handleDelete()}>
                        <RiDeleteBin6Fill />
                    </Button>
                </div>
            </td>

            {customerAssetUpdateModal &&
                <GenericModal content={<CreateUpdateCustomerAssetForm customerAsset={customerAsset} creditRequestId={customerAsset.creditRequestId} customerId={customerAsset.customerId} />}
                    isOpen={customerAssetUpdateModal} setOpen={setCustomerAssetUpdateModal} size='2xl' />}
        </tr>
    )
}

export default CustomerAssetRow