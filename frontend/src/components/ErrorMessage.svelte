<script>
    // Define props for the component
    export let message;
</script>

<div class="error-message">
    <p>{message}</p>
</div>

<style>
    /* Style the error message as needed */
    .error-message {
        background-color: #ff6b6b;
        color: white;
        padding: 10px;
        border-radius: 4px;
        border: 1px solid #ff1e1e;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }
</style>