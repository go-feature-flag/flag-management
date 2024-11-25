import { useDocumentTitle } from "../../../hooks/documentTitle.ts";

export const HomePage = () => {
  useDocumentTitle("Welcome");

  return (
    <>
      <div className="p-3 md:p-6 lg:p-10">
        <section>
          TODO:
          <ul>
            <li>API Key creation</li>
            <li>Exporter configuration</li>
            <li>Onboarding page</li>
          </ul>
        </section>
      </div>
    </>
  );
};
