import { useState } from 'react';
import { FaSearch, FaQuestionCircle, FaUserCircle, FaCloudUploadAlt} from 'react-icons/fa';
import { MdDocumentScanner, MdOutlineSchema, MdTextFields, MdSettings, MdOutlineStraighten } from 'react-icons/md';

const SidebarItem = ({ icon, label }) => (
  <div className="flex items-center gap-3 p-4 border-b hover:bg-gray-100 cursor-pointer">
    {icon}
    <span>{label}</span>
  </div>
);

export default function Home() {
  const [dragActive, setDragActive] = useState(false);

  return (
    <div className="flex h-screen font-sans">
      {/* Сайдбар */}
      <div className="w-64 bg-white border-r flex flex-col flex-1">
        <button>
          <SidebarItem icon={<MdDocumentScanner size={40} />} label="Свойства документа" />
        </button>

        <button>
          <SidebarItem icon={<MdOutlineSchema size={40} />} label="Структура документа" />
        </button>

        <button>
          <SidebarItem icon={<MdTextFields size={40} />} label="Текстовые метрики" />
        </button>

        <button>
          <SidebarItem icon={<MdSettings size={40} />} label="Параметры экспертизы" />
        </button>

        <button>
          <SidebarItem icon={<MdOutlineStraighten size={40} />} label="Размер документа" />
        </button>
      </div>

      {/* Контент */}
      <div className="flex-1 flex flex-col">
        {/* Верхняя панель */}
        <div className="flex items-center justify-between px-6 py-3 border-b bg-white">
          <div className="flex items-center gap-6">
            <button className="text-sm font-medium">Загрузить файл</button>
            <button className="text-sm font-medium">Библиотека файлов</button>
          </div>
          <div className="flex items-center gap-4">
            <div className="relative">
              <input
                type="text"
                placeholder="Поиск"
                className="border rounded pl-8 pr-2 py-1 text-sm"
              />
              <FaSearch className="absolute left-2 top-1.5 text-gray-500" />
            </div>
            <FaQuestionCircle size={20} />
            <FaCloudUploadAlt size={20} />
            <FaUserCircle size={22} />
          </div>
        </div>

        {/* Основной блок */}
        <div className="flex-1 flex items-center justify-center bg-gray-50">
          <div
            className={`w-80 h-48 border-2 border-dashed rounded-lg flex flex-col items-center justify-center ${
              dragActive ? 'border-blue-400 bg-blue-50' : 'border-gray-300'
            }`}
            onDragEnter={() => setDragActive(true)}
            onDragLeave={() => setDragActive(false)}
            onDrop={() => setDragActive(false)}
          >
            <p className="mb-2">Перетащите файл сюда</p>
            <button className="px-4 py-2 bg-gray-200 rounded-full hover:bg-gray-300">Выберите файл</button>
          </div>
        </div>
      </div>
    </div>
  );
}
