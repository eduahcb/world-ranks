const sortSelectEl = document.getElementById('sort');
const regionCheckboxesEl = document.querySelectorAll('input[name="region"]');
const statusRadioButtons = document.querySelectorAll('input[name="status"]');
const searchInputEl = document.getElementById('search');

const TIMEOUT = 700;


function debouce(fn, delay) {
  let timerId = null;

  return function(...args) {
    clearTimeout(timerId);
    timerId = setTimeout(() => fn.apply(this, args), delay);
  }
}

function updateUrlParams(params) {
  window.location.search = params.toString();
}

function buildParams() {
  const params = new URLSearchParams(window.location.search);

  if (sortSelectEl && sortSelectEl.value) {
    params.set('sort', sortSelectEl.value);
  }

  params.delete('region')
  regionCheckboxesEl.forEach((checkbox) => {
    if (checkbox.checked) {
      params.append('region', checkbox.value);
    }
  });

  let checkedStatus = null;

  statusRadioButtons.forEach((radio) => {
    if (radio.checked) {
      checkedStatus = radio;
    }
  })

  if (checkedStatus) {
    params.set('status', checkedStatus.value);
  }

  if (searchInputEl && searchInputEl.value) {
    params.set('search', searchInputEl.value);
  }
  else {
    params.delete('search');
  }

  return params;
}

function handleChange() {
  const params = buildParams();
  updateUrlParams(params);
}

sortSelectEl?.addEventListener('change', handleChange);

regionCheckboxesEl.forEach((checkbox) => {
  checkbox.addEventListener('change', handleChange);
});

statusRadioButtons.forEach((radio) => {
  radio.addEventListener('change', handleChange);
});


if (searchInputEl) {
  searchInputEl.addEventListener('keyup', debouce(handleChange, TIMEOUT));

  searchInputEl.addEventListener('search', (event) => {
    if (event.target.value === '') {
      handleChange();
    }
  });
}

window.addEventListener('DOMContentLoaded', () => {
  const params = new URLSearchParams(window.location.search);
  const search = params.get('search');

  if (search && searchInputEl) {
    searchInputEl.focus();

    const len = searchInputEl.value.length;
    searchInputEl.setSelectionRange(len, len);
  }
})
