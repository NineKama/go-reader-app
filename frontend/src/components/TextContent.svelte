<script>
    // Define props for the component
    export let textContent;

    // Define the number of lines to display per page
    let linesPerPage = 50;

    // State variable for font size
    let fontSize = 16; // Initial font size in pixels

    // Calculate the total number of pages
    let totalPages = Math.ceil(textContent.split('\n').length / linesPerPage);

    // Initialize the current page
    let currentPage = 1;

    // Function to go to the next page
    function nextPage() {
        if (currentPage < totalPages) {
            currentPage += 1;
        }
    }

    // Function to increase font size
    function increaseFontSize() {
        fontSize += 2; // You can adjust the increment as needed
    }

    // Function to decrease font size
    function decreaseFontSize() {
        fontSize -= 2; // You can adjust the decrement as needed
    }

    // Function to go to the previous page
    function prevPage() {
        if (currentPage > 1) {
            currentPage -= 1;
        }
    }
</script>

<main>
    <div class="text-content">
        {#if textContent}
            <!-- Display content for the current page -->
            {#if totalPages === 1}
                <!-- If there's only one page, display all content -->
                <pre style="font-size: {fontSize}px;">{textContent}</pre>
            {:else}
                <pre style="font-size: {fontSize}px;">
                    {#each textContent.split('\n').slice((currentPage - 1) * linesPerPage, currentPage * linesPerPage) as line}
                        {line}<br>
                    {/each}
                </pre>
                <div class="pagination">
                    {#if currentPage !== 1}
                        <button on:click={prevPage}>Previous</button>
                    {/if}
                    <span>Page {currentPage} of {totalPages}</span>
                    {#if currentPage !== totalPages}
                        <button on:click={nextPage}>Next </button>
                    {/if}
                </div>
            {/if}
        {:else}
            <p>No text content available.</p>
        {/if}
        <!-- Button to increase font size -->
        <button on:click={increaseFontSize}>Increase Font Size</button>

        <!-- Button to decrease font size -->
        <button on:click={decreaseFontSize}>Decrease Font Size</button>
    </div>
</main>

<style>
    /* Style the pagination buttons and other styles as needed */
    .pagination {
        margin-top: 10px;
    }

    /* Style the buttons and other styles as needed */
    .text-content {
        padding: 20px;
        background-color: #fff;
        border: 1px solid #ddd;
        border-radius: 4px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        color: #000; /* Set the text color to black */
    }

    button {
        margin: 10px;
        padding: 5px 10px;
        background-color: #007bff;
        color: #fff;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }

    button:disabled {
        background-color: #ccc;
        cursor: not-allowed;
    }
</style>