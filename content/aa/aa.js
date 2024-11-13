
// Function to extract tags from the summary text
function extractTags(summaryText) {
    const tagPattern = /\(([^)]+)\)/g;
    const matches = [];
    let match;
    while ((match = tagPattern.exec(summaryText)) !== null) {
        matches.push(...match[1].split(' '));
    }
    return matches.length > 0 ? matches : ["UNK"]; // Insert "UNK" if no tags found
}

// Function to assign data-tags to each <details> element based on the summary
function assignDataTags() {
    const detailsElements = document.querySelectorAll('details'); // Target all <details> elements
    detailsElements.forEach(details => {
        if (!details.classList.contains('details-item')) {
            details.classList.add('details-item'); // Add class if missing for consistency
        }
        const summaryText = details.querySelector('summary').textContent;
        const tags = extractTags(summaryText);
        details.setAttribute('data-tags', tags.join(' ')); // Set data-tags
    });
}

// Function to filter details based on selected filters
function filterDetails() {
    const selectedFilters = Array.from(document.querySelectorAll('.filter-checkbox:checked'))
        .map(checkbox => checkbox.value);

    const detailsItems = document.querySelectorAll('.details-item');
    detailsItems.forEach(item => {
        const tags = item.getAttribute('data-tags').split(' ');

        // Show all items if "ALL" is selected, otherwise apply specific filter logic
        const shouldShow = selectedFilters.includes("ALL") || tags.includes("ALL") || selectedFilters.some(filter => tags.includes(filter));
        item.style.display = shouldShow ? 'block' : 'none';
    });
}

// Debounced function for filtering (only runs after the user stops selecting filters)
let debounceTimeout;
function debounceFilter() {
    clearTimeout(debounceTimeout);
    debounceTimeout = setTimeout(filterDetails, 300); // Wait for 300ms after the last change
}

// Handle checkbox selection and enforce 3-filter limit
function handleCheckboxLimit(event) {
    const allCheckbox = document.querySelector('.filter-checkbox[value="ALL"]');
    const checkedCheckboxes = Array.from(document.querySelectorAll('.filter-checkbox:checked'));

    // Handle the case where "ALL" is selected
    if (event.target.value === 'ALL') {
        if (allCheckbox.checked) {
            document.querySelectorAll('.filter-checkbox:not([value="ALL"])').forEach(checkbox => {
                checkbox.checked = false;
            });
        }
    } else {
        allCheckbox.checked = false; // Uncheck "ALL" when any other box is selected
    }

    // Enforce 3-checkbox limit
    if (checkedCheckboxes.length > 3) {
        checkedCheckboxes[3].checked = false;
    }

    // Apply the filtering after checkbox change
    debounceFilter(); // Debounced call to filter
}

// Initialize the page
document.addEventListener('DOMContentLoaded', function() {
    assignDataTags(); // Assign tags to each details element based on summary text
    filterDetails(); // Initially apply the filtering

    // Add event listeners for checkbox changes
    document.querySelectorAll('.filter-checkbox').forEach(checkbox => {
        checkbox.addEventListener('change', handleCheckboxLimit);
    });
});