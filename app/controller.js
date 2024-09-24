var console = {
  log: function(msg) {
    print(`[JS Log]: ${msg}`);
  },
  error: function(msg) {
    print(`[JS Error]: ${msg}`);
  }
};

(async function() {
  if (!jsProgramPath || !hash) {
    console.error('Usage: controller.js <module-path> <hash>');
    return;
  }

  try {
    const program = await import(jsProgramPath);
    console.log('Running program:', jsProgramPath);

    if (typeof program.default === 'function') {
      program.default(hash);  
    } else {
      console.log('The specified program does not export a default function.');
    }

    console.log('Hash received:', hash);
  } catch (err) {
    console.error('Error running the program:', err.message);
  }
})();
