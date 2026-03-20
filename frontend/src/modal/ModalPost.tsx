import { toast } from "sonner";
import { useModalPost } from "../store/useModal.ts";
import { useUserPost } from "../store/usePostUser.ts";
import { Upload } from "lucide-react";

export default function ModalPost() {
  const { open, closeModal } = useModalPost((state) => state);
  const { form_post_user, SendData } = useUserPost((state) => state);

  const MAX_FILE_SIZE = 2 * 1024 * 1024; // 2 MB

  const isPdf = (file: File): boolean => {
    return (
      file.type === "application/pdf" ||
      file.name.toLowerCase().endsWith(".pdf")
    );
  };

  const handleChangeInput = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    useUserPost.setState((state) => ({
      form_post_user: {
        ...state.form_post_user,
        [name as keyof typeof state.form_post_user]: value,
      },
    }));
  };

  const handleDrop = async (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    const uploaded = e.dataTransfer.files[0];

    if (uploaded) {
      if (!isPdf(uploaded)) {
        toast("Solo se permiten archivos PDF");
        return;
      }

      if (uploaded.size > MAX_FILE_SIZE) {
        toast("El archivo no puede superar los 2 MB");
        return;
      }

      const fileData = await fileToArrayBuffer(uploaded);

      useUserPost.setState((state) => ({
        form_post_user: {
          ...state.form_post_user,
          file: fileData,
          fileName: uploaded.name,
        },
      }));
    }
  };

  const handleFile = async (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const uploaded = e.target.files[0];

      if (!isPdf(uploaded)) {
        toast("Solo se permiten archivos PDF");
        return;
      }

      if (uploaded.size > MAX_FILE_SIZE) {
        toast("El archivo no puede superar los 2 MB");
        return;
      }

      const fileData = await fileToArrayBuffer(uploaded);

      useUserPost.setState((state) => ({
        form_post_user: {
          ...state.form_post_user,
          file: fileData,
          fileName: uploaded.name,
        },
      }));
    }
  };

  const fileToArrayBuffer = (file: File): Promise<ArrayBuffer> => {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.onloadend = () => resolve(reader.result as ArrayBuffer);
      reader.onerror = reject;
      reader.readAsArrayBuffer(file);
    });
  };

  if (!open) return null;

  return (
    <div className="fixed inset-0 bg-black/60 flex items-center justify-center z-50">
      <div className="bg-gray-800 rounded-xl w-full max-w-md p-6 shadow-lg">
        <h2 className="text-lg font-semibold mb-4 text-gray-200">
          Nuevo registro
        </h2>

        <div className="mb-4">
          <label className="text-sm text-gray-400">Descripción</label>
          <textarea
            rows={3}
            placeholder="Descripcion"
            name="descripcion"
            value={form_post_user.descripcion}
            onChange={handleChangeInput}
            className="w-full mt-1 bg-gray-900 border border-gray-700 rounded-lg p-2 focus:outline-none focus:border-purple-500"
          />
        </div>

        <div
          onDrop={handleDrop}
          onDragOver={(e) => e.preventDefault()}
          className="border-2 border-dashed border-gray-600 rounded-lg p-6 text-center hover:border-purple-500 transition cursor-pointer mb-4"
        >
          <Upload className="mx-auto mb-2 text-purple-400" />
          <p className="text-sm text-gray-400">Arrastra un archivo aquí</p>
          <p className="text-xs text-gray-500">o haz click para subir</p>

          <input
            type="file"
            onChange={handleFile}
            className="hidden"
            id="fileUpload"
          />
          <label
            htmlFor="fileUpload"
            className="cursor-pointer block mt-2 text-purple-400 text-sm"
          >
            Seleccionar archivo
          </label>

          {form_post_user.fileName && (
            <p className="text-xs mt-2 text-green-400">
              Archivo cargado: {form_post_user.fileName}
            </p>
          )}
        </div>

        <div className="flex justify-end gap-3">
          <button
            onClick={closeModal}
            className="px-4 py-2 rounded-lg bg-gray-700 hover:bg-gray-600"
          >
            Cancelar
          </button>
          <button
            onClick={SendData}
            className="px-4 py-2 rounded-lg bg-purple-500 hover:bg-purple-600"
          >
            Guardar
          </button>
        </div>
      </div>
    </div>
  );
}
