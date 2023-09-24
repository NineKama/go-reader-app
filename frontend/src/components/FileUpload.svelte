<script>
    // Define props and variables for the component
    export let onFileUpload;

    // Function to handle file selection
    function handleFileChange(event) {
        const selectedFile = event.target.files[0];

        if (selectedFile) {
            // Call the parent component's callback function with the selected file
            onFileUpload(selectedFile);
        }
    }
</script>

<label for="fileInput">Upload a File:</label>
<input type="file" id="fileInput" accept=".txt,.pdf" on:change={handleFileChange} />

<style>
    /* Add styling for the file input and label as needed */
</style>