module.exports = {
    generateNotes: async (pluginConfig, context) => {
      return context.nextRelease.notes;
    },
    verifyConditions: async () => {
      console.log("✅ CalVer Plugin: Conditions verified.");
    },
    prepare: async (pluginConfig, context) => {
      const { lastRelease, nextRelease } = context;
      const now = new Date();
      const year = now.getFullYear().toString().slice(-2); // 2025 → 25
      const month = now.getMonth() + 1;
      let patch = 0;
  
      if (lastRelease && lastRelease.version) {
        const match = lastRelease.version.match(/^(\d{2})\.(\d+)\.(\d+)$/);
        if (match) {
          patch = parseInt(match[3], 10);
        }
      }
  
      patch += 1;
  
      nextRelease.version = `${year}.${month}.${patch}`;
      console.log(`🚀 Next Release Version: ${nextRelease.version}`);
    }
  };
  