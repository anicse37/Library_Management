const form = document.getElementById('bookForm');
const booksTableBody = document.getElementById('booksTableBody');

async function fetchBooks() {
    const res = await fetch('/books');
    const data = await res.json(); // <- Contains `book` array
    booksTableBody.innerHTML = '';
    data.book.forEach(book => {
        booksTableBody.innerHTML += `
          <tr>
            <td>${book.id}</td>
            <td>${book.name}</td>
            <td>${book.author}</td>
            <td>${book.description}</td>
            <td>${book.year}</td>
            <td>${book.available ? 'Yes' : 'No'}</td>
          </tr>
        `;
    });
}

form.addEventListener('submit', async (e) => {
    e.preventDefault();
    const book = {
        name: document.getElementById('name').value,
        author: document.getElementById('author').value,
        description: document.getElementById('description').value,
        year: parseInt(document.getElementById('year').value)
    };

    const res = await fetch('/books', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(book)
    });

    if (res.ok) {
        form.reset();
        fetchBooks();
    } else {
        alert('Failed to add book');
    }
});

window.onload = fetchBooks;
