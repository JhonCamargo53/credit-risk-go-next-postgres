import { create } from "zustand";
import { PageData } from "@/types/table";
import { DocumentType } from "@/types/documentType";
import { fetchDocumentTypes } from "@/services/documentTypeService";

interface DocumentTypeStore {
    documentTypes: DocumentType[];
    allDocumentTypes: DocumentType[];
    selectedDocumentType: DocumentType | null;

    pageData: PageData;

    loading: {
        fetching: boolean;
    };

    error: string | null;

    resetStatus: () => void;

    handlePagination: (pageData: PageData) => void;

    fetchDocumentTypes: () => Promise<void>;
}

export const useDocumentTypeStore = create<DocumentTypeStore>((set, get) => ({
    documentTypes: [],
    allDocumentTypes: [],
    selectedDocumentType: null,

    pageData: {
        currentPage: 1,
        limit: 10,
        total: 0,
        lastPage: 1,
    },

    loading: {
        fetching: false,
    },

    error: null,

    resetStatus: () => set({ error: null }),

    handlePagination: (pageData: PageData) => {
        const all = get().allDocumentTypes;
        const { currentPage, limit } = pageData;

        const start = (currentPage - 1) * limit;
        const end = start + limit;

        const paginated = all.slice(start, end);

        set({
            documentTypes: paginated,
            pageData: {
                currentPage,
                limit,
                total: all.length,
                lastPage: Math.ceil(all.length / limit),
            },
        });
    },

    fetchDocumentTypes: async () => {
        try {
            set((state) => ({
                loading: { ...state.loading, fetching: true },
                error: null,
            }));

            const data = await fetchDocumentTypes();

            set({ allDocumentTypes: data });

            get().handlePagination(get().pageData);
        } catch (err: any) {
            set({ error: err.message || "Error al obtener tipos de documento" });
            throw err;
        } finally {
            set((state) => ({
                loading: { ...state.loading, fetching: false },
            }));
        }
    },
}));
