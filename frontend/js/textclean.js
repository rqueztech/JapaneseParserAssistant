const frequencymap = new Map();
const meaningmap = new Map();

document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('textForm');
    const userInput = document.getElementById('userInput');
    const outputContainer = document.getElementById('outputContainer');

    form.addEventListener('submit', function(event) {
        event.preventDefault(); // Prevent form submission

        // Get the value from the textarea
        let inputText = userInput.value.trim(); // Use let instead of const
        
        // Eliminate leading spaces
        const eliminateLeadingSpaces = /^\s+/gm;
        inputText = inputText.replace(eliminateLeadingSpaces, '');

        // Eliminate unwanted words
        const eliminateUnwantedWords = /^(?![\u3040-\u309f\u30ao-\u30ff\u4e00-\u9fff]).*$/gm;
        inputText = inputText.replace(eliminateUnwantedWords, '');

        // Split into tokenized words (lines)
        const tokenizedWords = inputText.split('\n');

        // Create new elements for the output
        tokenizedWords.forEach(line => {

            if (line.trim()) { // Check if the line is not empty
                const removenonkanji = /[^\u4e00-\u9fff]/g;
                let kanjionlycleaned = line.replace(removenonkanji, '');

                const encodedInputText = encodeURIComponent(line);
                const jishoLink = `<a href="https://www.jisho.org/search/${line}" target="_blank">${line}</a>`;

                const newTitle = document.createElement('div');
                const newOutputBox = document.createElement('div');

                newTitle.classList.add('titlelinks');
                newTitle.innerHTML = jishoLink;
                outputContainer.appendChild(newTitle);
                
                for (let kanji = 0; kanji < kanjionlycleaned.length; kanji++) {
                    const newOutputBox = document.createElement('div');
                    newOutputBox.classList.add('box');
                    newOutputBox.textContent = kanjionlycleaned[kanji];
                    outputContainer.appendChild(newOutputBox);
                }
            }
        });

        // Clear the textarea after creating output box (optional)
        userInput.value = '';
    });
});

