<template>
    <div class="alarm">
        <div class="content">
            <div class="title">
                <h2>Feedback</h2>
                <!-- <h1>Form</h1> -->
            </div>
            <div>
                <table>
                    <thead>
                        <tr>
                            <th>Link</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>
                                <a
                                    href="https://docs.google.com/forms/d/1WfT8gZwkk7gSiUyvRf5gMnucaZonKOft33cgRZkL6Zs/prefill">
                                    Customer Device Usage Feedback Form
                                </a>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
    <div class="alarm">
        <div class="content">
            <div class="title">
                <h2>Download</h2>
                <h1>File</h1>
            </div>
            <div>
                <table>
                    <thead>
                        <tr>
                            <th>Manual</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>
                                <a
                                    href="https://drive.usercontent.google.com/u/0/uc?id=1YXq6pe24OwX-vW-wML_38_jPvuEdaYOl&export=download">
                                    FCP_manual_A4_220240503
                                </a>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <a
                                    href="https://drive.usercontent.google.com/u/0/uc?id=1e_uLZBpZsHb1U8m4TViYvM9FSrfNyTG6&export=download">
                                    FCM_manual_A4_20240503
                                </a>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <a
                                    href="https://drive.usercontent.google.com/u/0/uc?id=1gAOwKHU53gYEgKiC2Z4OTq6PG8aLvAgF&export=download">
                                    FCM_EasyUseTutorial_A4_2202400503
                                </a>
                            </td>
                        </tr>

                    </tbody>
                </table>
            </div>
        </div>
    </div>


</template>

<script setup>
</script>
<style scoped>
.alarm {
    height: auto;
    width: 70%;
    grid-column-start: 1;
    grid-column-end: 5;
    background-color: rgb(26 26 26 / 30%);
    border-radius: 1.8rem;
    padding: 0.6rem 1.2rem;
    margin: 1rem auto 0 auto;
}

.title {
    align-content: center;
}

.alarm .content {
    display: grid;
    gap: 1rem;
    grid-template-columns: 0.3fr 1fr;
    width: 100%;
    height: 100%;
}

.alarm h1 {
    color: #E5E5E5;
    font-size: 5rem;
    text-align: center;
    /* line-height: 9.3rem; */
    white-space: nowrap;
}

.alarm h2 {
    color: #E5E5E5;
    font-size: 2.9rem;
    text-align: center;
}

table {
    width: 100%;
    height: 95%;
    border-collapse: collapse;
}

table thead tr {
    height: 0.25rem;
}

table th {
    font-size: 1.5rem;
    color: #E5E5E5;
    vertical-align: top;
}

table tbody {
    background: rgb(255 255 255 / 70%);
}

table tbody td {
    border-bottom: 0.06rem solid #989898;
    text-align: center;
    height: 2.5rem;
    color: rgb(187, 23, 23);
    font-weight: 600;
}

table tbody tr:first-child td:first-child {
    border-top-left-radius: 1.25rem;
}

table tbody tr:first-child td:last-child {
    border-top-right-radius: 1.25rem;
}

table tbody tr:last-child td:first-child {
    border-bottom-left-radius: 1.25rem;
}

table tbody tr:last-child td:last-child {
    border-bottom-right-radius: 1.25rem;
}

table tbody tr:last-child td {
    border: none;
}

@media (max-width: 1024px) {
    .alarm .content {
        grid-template-columns: 1fr;
        grid-template-rows: 0.3fr 1fr;
    }

    .alarm .title {
        margin-bottom: 1rem;
    }
}
</style>