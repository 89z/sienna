'use strict';

function setBool(key, val) {
   if (Services.prefs.getPrefType(key) == Services.prefs.PREF_INVALID) {
      console.log('INVALID', key);
   }
   Services.prefs.setBoolPref(key, val);
   if (! Services.prefs.prefHasUserValue(key)) {
      console.log('DEFAULT', key);
   }
}

function setInt(key, val) {
   const n = Services.prefs.getPrefType(key);
   if (n == Services.prefs.PREF_INVALID) {
      console.log('INVALID', key);
   }
   Services.prefs.setIntPref(key, val);
   const b = Services.prefs.prefHasUserValue(key);
   if (! b) {
      console.log('DEFAULT', key);
   }
}

Services.prefs.resetUserPrefs();
// always ask me where to save files
setInt('browser.download.folderList', 0);
// disable new tab page
setBool('browser.newtabpage.enabled', false);
// do not provide search suggestions
setBool('browser.search.suggest.enabled', false);
// disable default browser nag
setBool('browser.shell.checkDefaultBrowser', false);
// show windows and tabs from last time
setInt('browser.startup.page', 3);
// disable delay hiding mute tab
setInt('browser.tabs.delayHidingAudioPlayingIconMS', 0);
// title bar
setBool('browser.tabs.drawInTitlebar', false);
// jumplist setting
setBool('browser.taskbar.lists.enabled', false);
// disable URL autocomplete
setBool('browser.urlbar.autoFill', false);
// switch to tab
setBool('browser.urlbar.suggest.openpage', false);
// browser console
setBool('devtools.chrome.enabled', true);
// disable notifications
setBool('dom.webnotifications.enabled', false);
// fuck you pocket piece of shit
setBool('extensions.pocket.enabled', false);
// fix default shitty jerky ass scrolling
setInt('general.smoothScroll.mouseWheel.durationMaxMS', 400);
setInt('general.smoothScroll.mouseWheel.durationMinMS', 200);
// allow autoplay
setInt('media.autoplay.default', 0);
// allow autoplay
setBool('media.block-autoplay-until-in-foreground', false);
// youtube
setInt('network.cookie.cookieBehavior', 1);
// youtube
setBool('privacy.trackingprotection.pbmode.enabled', false);
// github bookmarklets
setBool('security.csp.enable', false);
// remember passwords
setBool('signon.rememberSignons', false);
// disable crapbar in new tab, until clicked
setInt('ui.prefersReducedMotion', 1);
