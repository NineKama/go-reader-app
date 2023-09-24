<script>
  import { onMount } from 'svelte';

  import FileUpload from './components/FileUpload.svelte';
  import TextContent from './components/TextContent.svelte';
  import ErrorMessage from './components/ErrorMessage.svelte';

  // State variables
  let selectedFile = null;
  let textContent = null;
  let pdfContent = null;
  let errorMessage = null;

  // Function to handle file upload
  function handleFileUpload(file) {
    // Reset previous content and error message
    textContent = null;
    pdfContent = null;
    errorMessage = null;

    // Set the selected file
    selectedFile = file;

    // Determine file type and process accordingly
    if (file.type === 'text/plain') {
      // Handle text file
      handleTextFile(file);
    } else if (file.type === 'application/pdf') {
      // Handle PDF file
      handlePDFFile(file);
    } else {
      // Display error for unsupported file type
      errorMessage = 'Unsupported file type. Please upload a text or PDF file.';
    }
  }

  // Function to handle text file
  function handleTextFile(file) {
    // Read the text content from the file
    const reader = new FileReader();
    reader.onload = (event) => {
      textContent = event.target.result;
    };
    reader.readAsText(file);
  }

  // Function to handle PDF file (you can integrate your PDF rendering logic here)
  function handlePDFFile(file) {
    // Display a message or integrate a PDF viewer library
    // pdfContent = 'PDF content goes here';
    errorMessage = 'PDF file handling is not implemented in this example.';
  }

  // Clear selected file
  function clearFile() {
    selectedFile = null;
    textContent = null;
    pdfContent = null;
    errorMessage = null;
  }

  // Lifecycle hook to clear selected file on component load
  onMount(clearFile);
</script>

<main>
  <!-- Display the selected file content or error message -->
  {#if errorMessage}
    <ErrorMessage message={errorMessage} />
  {:else if selectedFile}
    {#if textContent}
      <TextContent textContent={textContent} />
    <!--{:else if pdfContent}-->
    <!--  <PDFViewer pdfContent={pdfContent} />-->
    {/if}
  {:else}
    <FileUpload onFileUpload={handleFileUpload} />
  {/if}
</main>